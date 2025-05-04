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
	senderName: pipe(string(), maxLength(60), nonEmpty()),
	senderPhone: pipe(string(), maxLength(60), nonEmpty()),
	senderEmail: pipe(string(), email(), nonEmpty()),
	senderStreet: pipe(string(), maxLength(255), nonEmpty()),
	senderCity: pipe(string(), maxLength(255), nonEmpty()),
	senderZip: pipe(string(), maxLength(255), nonEmpty()),
	senderCountry: pipe(string(), maxLength(255), nonEmpty()),
	billingName: pipe(string(), maxLength(60), nonEmpty()),
	billingPhone: pipe(string(), maxLength(60), nonEmpty()),
	billingEmail: pipe(string(), email(), nonEmpty()),
	billingStreet: pipe(string(), maxLength(255), nonEmpty()),
	billingCity: pipe(string(), maxLength(255), nonEmpty()),
	billingZip: pipe(string(), maxLength(255), nonEmpty()),
	billingCountry: pipe(string(), maxLength(255), nonEmpty()),
	model: optional(string()),
	serial: optional(string()),
	categories: pipe(array(string()), minLength(1)),
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
