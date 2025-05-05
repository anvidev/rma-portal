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

const contact = object({
	name: pipe(string(), minLength(3), maxLength(60), nonEmpty()),
	phone: pipe(string(), maxLength(60), nonEmpty()),
	email: pipe(string(), email(), nonEmpty()),
	street: pipe(string(), maxLength(255), nonEmpty()),
	city: pipe(string(), maxLength(255), nonEmpty()),
	zip: pipe(string(), maxLength(255), nonEmpty()),
	country: pipe(string(), maxLength(255), nonEmpty())
})

const schema = object({
	sender: contact,
	billing: contact,
	model: optional(string()),
	serial: optional(string()),
	categories: pipe(array(string()), minLength(1)),
	issue: pipe(string(), nonEmpty(), minLength(50), maxLength(500))
})

export const load: PageServerLoad = async () => {
	return { form: await superValidate(valibot(schema)) }
}

export const actions: Actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, valibot(schema))
		console.log(form)

		if (!form.valid) return fail(400, { form })

		// insert into db

		return message(form, 'form submitted')
	}
}
