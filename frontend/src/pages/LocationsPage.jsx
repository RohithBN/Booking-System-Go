import { useEffect, useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { apiGet } from '../api'

function LocationsPage() {
  const navigate = useNavigate()
  const [locations, setLocations] = useState([])
  const [query, setQuery] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    const load = async () => {
      setLoading(true)
      setError('')
      try {
        const data = await apiGet('/locations')
        setLocations(data)
      } catch (err) {
        setError(err.message)
      } finally {
        setLoading(false)
      }
    }
    load()
  }, [])

  const filtered = useMemo(
    () =>
      locations.filter((location) =>
        `${location.name} ${location.city}`.toLowerCase().includes(query.trim().toLowerCase()),
      ),
    [locations, query],
  )

  return (
    <main className="page">
      <section className="card">
        <header className="cardHeader">
          <p className="kicker">BookMyShow-style</p>
          <h1>Select Your Location</h1>
          <p className="subtitle">Search and pick a city to discover theatres and shows</p>
        </header>

        <div className="catalogBlock">
          <input
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            placeholder="Search location"
          />

          <div className="scrollList lg">
            {loading && <p className="meta">Loading locations...</p>}
            {!loading && filtered.length === 0 && <p className="meta">No locations found</p>}
            {filtered.map((location) => (
              <button
                key={location.id}
                type="button"
                className="listItemBtn"
                onClick={() => navigate(`/locations/${location.id}/theatres`, { state: { location } })}
              >
                <span>{location.name}</span>
                <small>{location.city}</small>
              </button>
            ))}
          </div>
        </div>

        {error && (
          <div className="statusArea">
            <p className="err">{error}</p>
          </div>
        )}
      </section>
    </main>
  )
}

export default LocationsPage
