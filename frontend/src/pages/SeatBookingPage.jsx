import { useCallback, useEffect, useMemo, useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { apiGet, confirmSeat, holdSeat } from '../api'

const ROWS = ['A', 'B', 'C', 'D', 'E', 'F']
const SEATS_PER_ROW = 10
const POLL_INTERVAL_MS = 2000

const buildDefaultSeatState = () => {
  const state = {}
  for (const row of ROWS) {
    for (let index = 1; index <= SEATS_PER_ROW; index += 1) {
      state[`${row}${index}`] = 'available'
    }
  }
  return state
}

function SeatBookingPage() {
  const navigate = useNavigate()
  const routeLocation = useLocation()
  const selectedShow = routeLocation.state?.show
  const selectedTheatre = routeLocation.state?.theatre
  const bookingMovieId = routeLocation.state?.bookingMovieId

  const [seatState, setSeatState] = useState(buildDefaultSeatState)
  const [selectedSeat, setSelectedSeat] = useState('')
  const [userId, setUserId] = useState('')
  const [holdId, setHoldId] = useState('')
  const [heldSeat, setHeldSeat] = useState('')
  const [booking, setBooking] = useState(null)
  const [isSyncing, setIsSyncing] = useState(false)
  const [message, setMessage] = useState('')
  const [error, setError] = useState('')
  const [isHolding, setIsHolding] = useState(false)
  const [isConfirming, setIsConfirming] = useState(false)

  const seatRows = useMemo(
    () =>
      ROWS.map((row) => ({
        row,
        seats: Array.from({ length: SEATS_PER_ROW }, (_, index) => `${row}${index + 1}`),
      })),
    [],
  )

  const syncSeats = useCallback(async () => {
    if (!bookingMovieId) {
      return
    }

    setIsSyncing(true)
    try {
      const bookings = await apiGet(`/bookings?movieId=${encodeURIComponent(bookingMovieId)}`)
      const nextState = buildDefaultSeatState()
      for (const item of bookings) {
        if (!item.SeatId || !nextState[item.SeatId]) {
          continue
        }
        if (item.Status === 'booked') {
          nextState[item.SeatId] = 'booked'
        } else if (item.Status === 'held') {
          nextState[item.SeatId] = 'held'
        }
      }
      setSeatState(nextState)
    } catch (err) {
      setError(err.message)
    } finally {
      setIsSyncing(false)
    }
  }, [bookingMovieId])

  useEffect(() => {
    if (!bookingMovieId) {
      return
    }

    syncSeats()
    const intervalId = window.setInterval(syncSeats, POLL_INTERVAL_MS)
    return () => window.clearInterval(intervalId)
  }, [bookingMovieId, syncSeats])

  const onHoldSeat = async (event) => {
    event.preventDefault()
    setError('')
    setMessage('')
    setBooking(null)

    if (!bookingMovieId) {
      setError('Show not selected. Go back and choose a show.')
      return
    }

    if (!selectedSeat) {
      setError('Select a seat first')
      return
    }

    if (seatState[selectedSeat] === 'booked') {
      setError('This seat is already booked')
      return
    }

    setIsHolding(true)
    try {
      const data = await holdSeat({
        MovieId: bookingMovieId,
        SeatId: selectedSeat,
        UserId: userId,
      })
      setSeatState((prev) => ({ ...prev, [selectedSeat]: 'held' }))
      setHeldSeat(selectedSeat)
      setHoldId(data.id)
      setMessage(data.message || 'Seat held successfully')
    } catch (err) {
      setError(err.message)
      if (err.message.toLowerCase().includes('already held')) {
        setSeatState((prev) => ({ ...prev, [selectedSeat]: 'held' }))
      }
    } finally {
      setIsHolding(false)
    }
  }

  const onConfirm = async (event) => {
    event.preventDefault()
    setError('')
    setMessage('')
    setBooking(null)

    if (!holdId) {
      setError('Hold a seat first to get booking ID')
      return
    }

    setIsConfirming(true)
    try {
      const data = await confirmSeat(holdId)
      if (data.SeatId) {
        setSeatState((prev) => ({ ...prev, [data.SeatId]: 'booked' }))
      } else if (heldSeat) {
        setSeatState((prev) => ({ ...prev, [heldSeat]: 'booked' }))
      }
      setBooking(data)
      setMessage('Booking confirmed successfully')
      setHoldId('')
      setHeldSeat('')
      setSelectedSeat('')
    } catch (err) {
      setError(err.message)
    } finally {
      setIsConfirming(false)
    }
  }

  return (
    <main className="page">
      <section className="card">
        <header className="cardHeader">
          <p className="kicker">Seat Selection</p>
          <h1>{selectedShow?.movie?.title || 'Show not selected'}</h1>
          <p className="subtitle">
            {selectedTheatre?.name || 'Theatre'} •{' '}
            {selectedShow?.startsAt
              ? new Date(selectedShow.startsAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
              : '--:--'}
          </p>
          <p className="subtitle small">Movie ID: {bookingMovieId || '-'} {isSyncing ? '• syncing seats...' : ''}</p>
        </header>

        <div className="seatCard">
          <div className="screen">SCREEN</div>
          <div className="seatMap">
            {seatRows.map((seatRow) => (
              <div className="seatRow" key={seatRow.row}>
                <span className="rowLabel">{seatRow.row}</span>
                <div className="seatRowGrid">
                  {seatRow.seats.map((seat) => {
                    const status = seatState[seat] || 'available'
                    const isSelected = seat === selectedSeat
                    return (
                      <button
                        key={seat}
                        type="button"
                        className={`seat ${status} ${isSelected ? 'selected' : ''}`}
                        onClick={() => setSelectedSeat(seat)}
                        disabled={status === 'booked'}
                        title={`Row ${seatRow.row}, Seat ${seat.slice(1)} (${status})`}
                      >
                        {seat.slice(1)}
                      </button>
                    )
                  })}
                </div>
              </div>
            ))}
          </div>

          <div className="legend">
            <span><i className="dot availableDot" /> Available</span>
            <span><i className="dot selectedDot" /> Selected</span>
            <span><i className="dot heldDot" /> On Hold</span>
            <span><i className="dot bookedDot" /> Booked</span>
          </div>
        </div>

        <div className="grid">
          <form className="panel" onSubmit={onHoldSeat}>
            <h2>Hold Selected Seat</h2>
            <p className="meta">Movie: {bookingMovieId || '-'}</p>
            <p className="meta">Seat: {selectedSeat || '-'}</p>
            <label>
              User ID
              <input
                value={userId}
                onChange={(event) => setUserId(event.target.value)}
                placeholder="user-42"
              />
            </label>
            <button type="submit" disabled={isHolding || !selectedSeat}>
              {isHolding ? 'Holding...' : 'Hold Seat'}
            </button>
          </form>

          <form className="panel" onSubmit={onConfirm}>
            <h2>Confirm Current Hold</h2>
            <p className="meta">Booking ID: {holdId || '-'}</p>
            <button type="submit" disabled={isConfirming || !holdId}>
              {isConfirming ? 'Confirming...' : 'Confirm Booking'}
            </button>
            {holdId && <p className="hint">Use this ID in API tests: {holdId}</p>}
          </form>
        </div>

        <div className="navActions">
          <button type="button" className="backBtn" onClick={() => navigate(-1)}>
            ← Back to shows
          </button>
          <button type="button" className="backBtn" onClick={() => navigate('/')}>
            Home
          </button>
        </div>

        {(message || error) && (
          <div className="statusArea">
            {message && <p className="ok">{message}</p>}
            {error && <p className="err">{error}</p>}
          </div>
        )}

        {booking && (
          <div className="result">
            <h3>Confirmed Booking</h3>
            <dl>
              <div><dt>ID</dt><dd>{booking.ID}</dd></div>
              <div><dt>Movie</dt><dd>{booking.MovieId}</dd></div>
              <div><dt>Seat</dt><dd>{booking.SeatId}</dd></div>
              <div><dt>Status</dt><dd>{booking.Status}</dd></div>
            </dl>
          </div>
        )}
      </section>
    </main>
  )
}

export default SeatBookingPage
