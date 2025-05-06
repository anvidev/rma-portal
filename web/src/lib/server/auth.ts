import type { RequestEvent } from '@sveltejs/kit'
import { API_URL } from './env'

export const validateAuth = async (event: RequestEvent) => {
	const { cookies } = event

	const token = cookies.get('token')

	if (!token) {
		return null
	}

	const response = await fetch(`${API_URL}/v1/auth/validate`, {
		method: 'get',
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	if (!response.ok) {
		cookies.set('token', '', {
			path: '/',
			secure: true,
			httpOnly: true,
			maxAge: -1,
		})
		return null
	}

	const { access_token, user } = await response.json()

	cookies.set('token', access_token, {
		path: '/',
		secure: true,
		httpOnly: true,
	})

	return user
}
