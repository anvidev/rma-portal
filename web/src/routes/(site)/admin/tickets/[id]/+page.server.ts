import { API_URL } from '$lib/server/env'
import type { TicketWithLogs } from '$lib/types'
import { fail, redirect, type Actions } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { maxLength, nonEmpty, object, pipe, string } from 'valibot'
import { message, superValidate } from 'sveltekit-superforms'
import { valibot } from 'sveltekit-superforms/adapters'

const createLogSchema = object({
	status: pipe(string(), nonEmpty()),
	external_comment: pipe(string(), nonEmpty(), maxLength(200)),
	internal_comment: pipe(string(), maxLength(200)),
})

export const load: PageServerLoad = async ({ cookies, params, fetch }) => {
	const id = params.id

	const emptyForm = await superValidate(valibot(createLogSchema))

	const response = await fetch(`${API_URL}/v1/admin/tickets/${id}`, {
		headers: {
			Authorization: `Bearer ${cookies.get('token')}`,
		},
	})

	const statusResponse = await fetch(`${API_URL}/v1/tickets/statuses`)

	if (!response.ok || !statusResponse.ok) {
		redirect(308, '/admin/tickets')
	}

	const { ticket } = await response.json()
	const { statuses } = await statusResponse.json()

	return {
		form: emptyForm,
		ticket: ticket as TicketWithLogs,
		statuses: statuses as string[],
	}
}

export const actions: Actions = {
	default: async ({ request, fetch, params, cookies }) => {
		const form = await superValidate(request, valibot(createLogSchema))

		if (!form.valid) {
			return fail(400, { form })
		}

		const response = await fetch(`${API_URL}/v1/admin/tickets/${params.id}/log`, {
			method: 'post',
			headers: {
				Authorization: `Bearer ${cookies.get('token')}`,
			},
			body: JSON.stringify(form.data),
		})

		if (!response.ok) {
			return fail(response.status, { form })
		}

		return message(form, 'Opdatering tilføjet til RMA')
	},
}
