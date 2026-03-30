import { useEffect, useMemo, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { apiGet } from '../api'

function TheatresShowsPage() {
  const navigate = useNavigate()
  const { locationId } = useParams()

  const [theatres, setTheatres] = useState([])
  const [selectedTheatre, setSelectedTheatre] = useState(null)
  const [shows, setShows] = useState([])
  const [loadingTheatres, setLoadingTheatres] = useState(false)
  const [loadingShows, setLoadingShows] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    const loadTheatres = async () => {
      setLoadingTheatres(true)
      setError('')
      try {
        const data = await apiGet(`/theatres?locationId=${locationId}`)
        setTheatres(data)
        setSelectedTheatre(data[0] || null)
      } catch (err) {
        setError(err.message)
      } finally {
        setLoadingTheatres(false)
      }
    }
    loadTheatres()
  }, [locationId])

  useEffect(() => {
    if (!selectedTheatre?.id) {
      setShows([])
      return
    }

    const loadShows = async () => {
      setLoadingShows(true)
      setError('')
      try {
        const data = await apiGet(`/shows?theatreId=${selectedTheatre.id}`)
        setShows(data)
      } catch (err) {
        setError(err.message)
      } finally {
        setLoadingShows(false)
      }
    }
    loadShows()
  }, [selectedTheatre])

  const groupedShows = useMemo(() => {
    const grouped = new Map()
    for (const show of shows) {
      const movieId = show.movie?.id || show.movieId
      if (!grouped.has(movieId)) {
        grouped.set(movieId, { movie: show.movie, entries: [] })
      }
      grouped.get(movieId).entries.push(show)
    }
    return Array.from(grouped.values())
  }, [shows])

  return (
    <main className="page">
      <section className="card">
        <header className="cardHeader">
          <p className="kicker">Step 2</p>
          <h1>Select Theatre & Show</h1>
          <p className="subtitle">Location #{locationId}</p>
        </header>

        <div className="catalogFlow two">
          <div className="catalogBlock">
            <p className="blockTitle">Theatres</p>
            <div className="scrollList lg">
              {loadingTheatres && <p className="meta">Loading theatres...</p>}
              {!loadingTheatres && theatres.length === 0 && <p className="meta">No theatres found</p>}
              {theatres.map((theatre) => (
                <button
                  key={theatre.id}
                  type="button"
                  className={`listItemBtn ${selectedTheatre?.id === theatre.id ? 'active' : ''}`}
                  onClick={() => setSelectedTheatre(theatre)}
                >
                  {theatre.name}
                </button>
              ))}
            </div>
          </div>

          <div className="catalogBlock">
            <p className="blockTitle">Movies & Shows</p>
            <div className="scrollList lg">
              {loadingShows && <p className="meta">Loading shows...</p>}
              {!loadingShows && groupedShows.length === 0 && <p className="meta">No shows found</p>}

              {groupedShows.map((group) => (
                <div key={group.movie?.id || group.movie?.title} className="showGroup">
                  <p className="showMovieTitle">{group.movie?.title || 'Movie'}</p>
                  <div className="showTimes">
                    {group.entries.map((show) => (
                      <button
                        key={show.id}
                        type="button"
                        className="showTimeBtn"
                        onClick={() =>
                          navigate('/seats', {
                            state: {
                              show,
                              theatre: selectedTheatre,
                              bookingMovieId: show.movie?.code || `movie-${show.movieId}`,
                            },
                          })
                        }
                      >
                        {new Date(show.startsAt).toLocaleTimeString([], {
                          hour: '2-digit',
                          minute: '2-digit',
                        })}
                      </button>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        <button type="button" className="backBtn" onClick={() => navigate('/')}>
          ← Back to locations
        </button>

        {error && (
          <div className="statusArea">
            <p className="err">{error}</p>
          </div>
        )}
      </section>
    </main>
  )
}

export default TheatresShowsPage
