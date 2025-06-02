import { error, fail, type Actions } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { ApiError } from '$lib/server/api'
import { z } from 'zod'

const createLogSchema = z.object({
	status: z.string(),
	external_comment: z.string().max(200),
	internal_comment: z.string().max(200),
})

export type NewTicketLog = z.infer<typeof createLogSchema>

export const load: PageServerLoad = async ({ cookies, params, locals }) => {
	const id = params.id

	const emptyForm = await superValidate(zod(createLogSchema))

	const token = cookies.get('token') ?? ''

	const [ticketData, err] = await locals.api.getTicket(token, id)
	if (err != null) {
		if (err instanceof ApiError) {
			return error(err.status, { message: err.message, requestId: err.requestID })
		} else {
			return error(500, { message: err.message })
		}
	}

	const { ticket } = ticketData

	const [statusesData, statusErr] = await locals.api.listStatuses()
	if (statusErr != null) {
		if (statusErr instanceof ApiError) {
			return error(statusErr.status, { message: statusErr.message, requestId: statusErr.requestID })
		} else {
			return error(500, statusErr.message)
		}
	}

	const { statuses } = statusesData

	return {
		form: emptyForm,
		ticket: ticket,
		statuses: statuses,
	}
}

export const actions: Actions = {
	default: async ({ request, params, cookies, setHeaders, locals }) => {
		const form = await superValidate(request, zod(createLogSchema))

		if (!form.valid) {
			return fail(400, { form })
		}

		const token = cookies.get('token') ?? ''
		const id = params.id ?? ''

		const [_logData, err] = await locals.api.createTicketLog(token, id, form.data)
		if (err != null) {
			if (err instanceof ApiError) {
				return fail(err.status, { form })
			} else {
				return fail(500, { form })
			}
		}

		setHeaders({
			'Clear-Site-Data': 'cache',
		})

		return message(form, 'Opdatering tilføjet til RMA')
	},
}
