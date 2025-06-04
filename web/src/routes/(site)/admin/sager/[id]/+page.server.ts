import { error, type Actions } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { fail, message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { ApiError } from '$lib/server/api'
import { z } from 'zod'

const createLogSchema = z.object({
	status: z.string(),
	external_comment: z.string().max(200),
	internal_comment: z.string().max(200),
	files: z
		.instanceof(File, { message: 'Foo' })
		.refine(f => f.size < 4_000_000_000, 'Filerne må ikke være større end 4MB')
		.array()
		.max(10, { message: 'Maks. 10 filer må uploades' })
		.optional(),
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

		const { files, ...newLog } = form.data

		if (!form.valid) {
			return fail(400, { form })
		}

		const token = cookies.get('token') ?? ''
		const id = params.id ?? ''

		const [logData, err] = await locals.api.createTicketLog(token, id, newLog)
		if (err != null) {
			if (err instanceof ApiError) {
				return fail(err.status, { form })
			} else {
				return fail(500, { form })
			}
		}

		if (files && files.length > 0) {
			const formData = new FormData()

			files.forEach(f => formData.append('files', f))

			const [_filesData, err] = await locals.api.createLogFiles(token, id, logData.log.id, formData)
			if (err != null) {
				console.error(`fil upload fejlede for RMA opdatering #${logData.log.id}: ${err.message}`)
				if (err instanceof ApiError) {
					return fail(err.status, { form })
				}
				return fail(500, { form })
			}
		}

		setHeaders({
			'Clear-Site-Data': 'cache',
		})

		return message(form, 'Opdatering tilføjet til RMA')
	},
}
