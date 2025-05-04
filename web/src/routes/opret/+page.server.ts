import type { Actions, PageServerLoad } from './$types'
import {
	array,
	email,
	maxLength,
	minLength,
	nonEmpty,
	object,
	optional,
	pipe,
	string
} from 'valibot'
import { superValidate, message } from 'sveltekit-superforms'
import { valibot } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

const schema = object({
	senderName: pipe(string(), maxLength(60)),
	senderAddress: pipe(string(), maxLength(255)),
	senderPhone: pipe(string(), maxLength(60)),
	senderEmail: pipe(string(), email()),
	billingName: pipe(string(), maxLength(60)),
	billingAddress: pipe(string(), maxLength(255)),
	billingPhone: pipe(string(), maxLength(60)),
	billingEmail: pipe(string(), email()),
	model: optional(string()),
	serial: optional(string()),
	categories: array(string()),
	issue: pipe(string(), nonEmpty(), minLength(50), maxLength(500))
})

export const load: PageServerLoad = async () => {
	return { form: await superValidate(valibot(schema)) }
}

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event.request, valibot(schema))
		console.log(form)

		if (!form.valid) return fail(400, { form })

		// insert into db

		return message(form, 'form submitted')
	}
}
