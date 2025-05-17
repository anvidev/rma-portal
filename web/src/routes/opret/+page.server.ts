import type { Actions, PageServerLoad } from './$types'
import { superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail, redirect } from '@sveltejs/kit'
import { API_URL } from '$lib/server/env'
import * as z from 'zod'
import type { Ticket } from '$lib/types'

const radio = z
	.enum(['yes', 'no', 'unknown', 'none'])
	.refine(val => val != 'none', { message: 'Vælg venligst en af mulighederne' })

const contactZod = z.object({
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
		.max(2, { message: 'Vælg en af mulighederne' }),
})

const schema = z.object({
	sender: contactZod,
	billing: contactZod,
	model: z.string().optional(),
	serial: z.string().optional(),
	categories: z.array(z.string()).min(1, { message: 'Vælg minimum én af mulighederne' }),
	issue: z
		.string()
		.nonempty({ message: 'Dette felt er påkrævet' })
		.min(20, { message: 'Må ikke være mindre en 20 karakterer' })
		.max(500, { message: 'Må ikke være mere end 500 karaterer' }),
	quote: radio.default('none'),
	warranty: radio.default('none'),
})

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(schema)),
	}
}

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const form = await superValidate(request, zod(schema))

		if (!form.valid) return fail(400, { form })

		const response = await fetch(`${API_URL}/v1/tickets`, {
			method: 'post',
			body: JSON.stringify(form.data),
		})

		if (!response.ok) {
			return fail(response.status, { form })
		}

		const data = (await response.json()) as { ticket: Ticket }

		redirect(303, `/admin/tickets/${data.ticket.id}`)
	},
}
