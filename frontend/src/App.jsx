import { Navigate, Route, Routes } from 'react-router-dom'
import './App.css'
import LocationsPage from './pages/LocationsPage'
import SeatBookingPage from './pages/SeatBookingPage'
import TheatresShowsPage from './pages/TheatresShowsPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<LocationsPage />} />
      <Route path="/locations/:locationId/theatres" element={<TheatresShowsPage />} />
      <Route path="/seats" element={<SeatBookingPage />} />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}

export default App
