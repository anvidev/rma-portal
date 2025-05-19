import { redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'
import { API_URL } from '$lib/server/env'

export const load: PageServerLoad = async evt => {
	const accessToken = evt.cookies.get('token')

	if (accessToken) {
		throw redirect(301, '/admin/tickets')
	}
}

export const actions: Actions = {
	default: async evt => {
		const form = await evt.request.formData()
		const email = form.get('email')
		const password = form.get('password')

		const response = await fetch(`${API_URL}/v1/auth/login`, {
			method: 'post',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ email, password }),
		})

		const data = await response.json()

		if (response.ok) {
			const accessToken = data.access_token
			evt.cookies.set('token', accessToken, {
				path: '/',
				secure: true,
				httpOnly: true,
				maxAge: 60 * 60 * 24 * 3,
			})

			const redirectUrl = evt.url.searchParams.get('redirect') ?? '/admin/tickets'

			throw redirect(301, redirectUrl)
		}
	},
}
