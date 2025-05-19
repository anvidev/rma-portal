import { redirect, type Actions } from '@sveltejs/kit'

export const actions: Actions = {
	logout: async ({ cookies }) => {
		cookies.delete('token', { path: '/' })
		throw redirect(303, '/log-ind')
	},
}
