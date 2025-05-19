import { validateAuth } from '$lib/server/auth'
import { redirect, type Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = await validateAuth(event)

	// if (event.url.pathname == '/') {
	// 	throw redirect(303, '/log-ind')
	// }

	if (event.url.pathname.startsWith('/admin')) {
		if (event.locals.user == null) {
			throw redirect(303, `/log-ind?redirect=${event.url.pathname}`)
		}
	}

	const response = resolve(event)
	return response
}
