export async function apiGet(path) {
  const response = await fetch(`/api${path}`)
  const data = await response.json()

  if (!response.ok) {
    throw new Error(data.error || 'Request failed')
  }

  return data
}

export async function holdSeat(payload) {
  const response = await fetch('/api/hold', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })
  const data = await response.json()

  if (!response.ok) {
    throw new Error(data.error || 'Failed to hold seat')
  }

  return data
}

export async function confirmSeat(bookingId) {
  const response = await fetch(`/api/confirm/${bookingId}`, {
    method: 'POST',
  })
  const data = await response.json()

  if (!response.ok) {
    throw new Error(data.error || 'Failed to confirm booking')
  }

  return data
}
