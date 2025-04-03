'use client'

import { useEffect, useState } from 'react'

export default function ProfilePage() {
  const [message, setMessage] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const token = localStorage.getItem('token')

    if (!token) {
      setMessage('Non autorisé. Veuillez vous connecter.')
      setLoading(false)
      return
    }

    fetch('http://localhost:8080/profile', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then(async (res) => {
        const data = await res.json()
        if (res.ok) {
          setMessage(data.message)
        } else {
          setMessage(data.message || 'Erreur')
        }
      })
      .catch(() => {
        setMessage('Erreur serveur')
      })
      .finally(() => {
        setLoading(false)
      })
  }, [])

  return (
    <main className="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div className="max-w-md w-full bg-white p-8 rounded-xl shadow">
        <h2 className="text-2xl font-bold mb-4 text-center">Mon espace</h2>
        {loading ? (
          <p className="text-center text-gray-600">Chargement...</p>
        ) : (
          <p className="text-center text-gray-800">{message}</p>
        )}
      </div>
    </main>
  )
}
