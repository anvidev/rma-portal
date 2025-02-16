import { redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async (evt) => {
	const accessToken = evt.cookies.get('token')

	if (accessToken) {
		throw redirect(301, '/')
	}
}

export const actions: Actions = {
	default: async (evt) => {
		const form = await evt.request.formData()
		const email = form.get('email')
		const password = form.get('password')

		const response = await fetch('http://localhost:8080/v1/auth/login', {
			method: 'post',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		})

		const data = await response.json()

		if (response.ok) {
			const accessToken = data.access_token
			evt.cookies.set('token', accessToken, {
				path: '/',
				secure: true,
				httpOnly: true
			})

			throw redirect(301, '/admin/tickets')
		}
	}
}
