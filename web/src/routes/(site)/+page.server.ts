import { fail, setError, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { z } from 'zod'
import type { Actions, PageServerLoad } from './$types'
import { redirect } from '@sveltejs/kit'
import { ApiError } from '$lib/server/api'

const schema = z.object({
	query: z
		.string()
		.regex(/^\d{5}-[A-Za-z0-9]{10}$/, 'Ugyldigt RMA sags ID. Tjek formatet og prøv igen.'),
})

export const load: PageServerLoad = async () => {
	return { form: await superValidate(zod(schema)) }
}

export const actions: Actions = {
	default: async ({ request, locals }) => {
		const form = await superValidate(request, zod(schema))

		if (!form.valid) {
			return fail(400, { form })
		}

		const uppercasedID = form.data.query.toUpperCase()

		const [_ticketData, err] = await locals.api.getPublicTicket(uppercasedID)
		if (err != null) {
			if (err instanceof ApiError) {
				let msg = 'Der gik noget galt på serveren.'
				if (err.status == 404) {
					msg = 'RMA sag blev ikke fundet'
				}
				return setError(form, 'query', msg)
			} else {
				return setError(form, 'query', 'Intern server fejl')
			}
		}

		return redirect(303, `/sager/${uppercasedID}`)
	},
}
