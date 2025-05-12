import type { Actions, PageServerLoad } from './$types'
import {
	array,
	boolean,
	email,
	enum_,
	literal,
	maxLength,
	minLength,
	nonEmpty,
	object,
	optional,
	pipe,
	string,
} from 'valibot'
import { superValidate, message } from 'sveltekit-superforms'
import { valibot } from 'sveltekit-superforms/adapters'
import { fail, redirect } from '@sveltejs/kit'
import { API_URL } from '$lib/server/env'

enum WantQuote {
	Yes,
	No,
}

const contact = object({
	name: pipe(
		string('Skal være tekst'),
		nonEmpty('Navn er påkrævet'),
		minLength(3, 'Navn skal være mindst 3 tegn'),
		maxLength(60, 'Navn må maks være 60 tegn'),
	),
	phone: pipe(
		string('Skal være tekst'),
		nonEmpty('Telefonnr. er påkrævet'),
		maxLength(60, 'Telefon må maks være 60 tegn'),
	),
	email: pipe(
		string('Skal være tekst'),
		nonEmpty('E-mail er påkrævet'),
		email('Indtast en gyldig email'),
	),
	street: pipe(
		string('Skal være tekst'),
		nonEmpty('Vejnavn er påkrævet'),
		maxLength(255, 'Vejnavn må maks være 255 tegn'),
	),
	city: pipe(
		string('Skal være tekst'),
		nonEmpty('By er påkrævet'),
		maxLength(255, 'By må maks være 255 tegn'),
	),
	zip: pipe(
		string('Skal være tekst'),
		nonEmpty('Postnr. er påkrævet'),
		maxLength(255, 'Postnr må maks være 255 tegn'),
	),
	country: pipe(
		string('Skal være tekst'),
		nonEmpty('Land er påkrævet'),
		maxLength(2, 'Land må maks være 2 tegn'),
	),
})

const schema = object({
	sender: contact,
	billing: contact,
	model: optional(string()),
	serial: optional(string()),
	categories: pipe(array(string()), minLength(1, 'Vælg mindst én kategori')),
	issue: pipe(
		string('Skal være tekst'),
		nonEmpty('Problem er påkrævet'),
		minLength(50, 'Beskrivelse skal være mindst 50 tegn'),
		maxLength(500, 'Beskrivelse må maks være 500 tegn'),
	),
	wantsQuote: enum_(WantQuote, 'Vælg om du ønsker tilbud ved pris over minimum'),
})

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(valibot(schema)),
	}
}

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const form = await superValidate(request, valibot(schema))

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

		// return message(form, `ticket id ${data.ticket.id} created`)
	},
}
