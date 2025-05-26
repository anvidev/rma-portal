import type { Actions, PageServerLoad } from './$types'
import { fail, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { redirect } from '@sveltejs/kit'
import { API_URL } from '$lib/server/env'
import * as z from 'zod'
import { ApiError } from '$lib/server/api'

const radio = z
	.enum(['yes', 'no', 'unknown', 'none'])
	.refine(val => val != 'none', { message: 'Vælg venligst en af mulighederne' })

const contactZod = z.object({
	company: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.min(3, { message: 'Må ikke være mindre end 3 karakterer' })
		.max(100, { message: 'Må ikke være mere end 100 karakterer' }),
	name: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.min(3, { message: 'Må ikke være mindre end 3 karakterer' })
		.max(60, { message: 'Må ikke være mere end 60 karakterer' }),
	phone: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.max(60, { message: 'Må ikke være mere end 60 karakterer' }),
	email: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.email({ message: 'Skal være en valid email' }),
	street: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.max(255, { message: 'Må ikke være mere end 255 karakterer' }),
	city: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.max(255, { message: 'Må ikke være mere end 255 karakterer' }),
	zip: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.max(255, { message: 'Må ikke være mere end 255 karakterer' }),
	country: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.max(2, { message: 'Vælg en af mulighederne' })
		.default('DK'),
})

const schema = z.object({
	sender: contactZod,
	billing: contactZod,
	model: z.string().optional(),
	serial_number: z.string().optional(),
	categories: z.array(z.string()).min(1, { message: 'Vælg minimum én af mulighederne' }),
	issue: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.min(20, { message: 'Må ikke være mindre en 20 karakterer' })
		.max(500, { message: 'Må ikke være mere end 500 karaterer' }),
	quote: radio.default('none'),
	warranty: radio.default('none'),
	files: z
		.instanceof(File, { message: 'Foo' })
		.refine(f => f.size < 4_000_000_000, 'Billederne må ikke være større end 4MB')
		.array()
		.max(4, { message: 'Maks. 4 billeder må uploades' })
		.optional(),
})

export type NewTicket = z.infer<typeof schema>

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(schema)),
	}
}

export const actions: Actions = {
	default: async ({ request, locals, setHeaders }) => {
		const form = await superValidate(request, zod(schema))
		if (!form.valid) return fail(400, { form })

		const ticketFormData = {
			sender: form.data.sender,
			billing: form.data.billing,
			model: form.data.model,
			serial_number: form.data.serial_number,
			categories: form.data.categories,
			issue: form.data.issue,
			quote: form.data.quote,
			warranty: form.data.warranty,
		}

		const [ticketData, err] = await locals.api.createTicket(ticketFormData)
		if (err != null) {
			if (err instanceof ApiError) {
				return fail(err.status, { form })
			}
			return fail(500, { form })
		}

		const { ticket } = ticketData

		if (form.data.files && form.data.files.length > 0) {
			const formData = new FormData()
			form.data.files.forEach(f => {
				formData.append('files', f)
			})

			const fileResponse = await fetch(`${API_URL}/v1/tickets/${ticket.id}/files`, {
				method: 'POST',
				body: formData,
			})

			if (!fileResponse.ok) {
				const { error } = await fileResponse.json()
				console.error(`fil upload fejlede for RMA #${ticket.id}: ${error ?? 'ukendt fejl'}`)
			}
		}

		const redirectUrl = locals.user ? `/admin/sager/${ticket.id}` : `sager/${ticket.id}/tak`

		if (locals.user) {
			setHeaders({
				'Clear-Site-Data': 'cache',
			})
		}

		redirect(303, redirectUrl)
	},
}
