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
	string,
} from 'valibot'
import { superValidate, message } from 'sveltekit-superforms'
import { valibot } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

const contact = object({
	name: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld navn'),
		minLength(3, 'Navn skal være mindst 3 tegn'),
		maxLength(60, 'Navn må maks være 60 tegn'),
	),
	phone: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld telefonnummer'),
		maxLength(60, 'Telefon må maks være 60 tegn'),
	),
	email: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld email'),
		email('Indtast en gyldig email'),
	),
	street: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld adresse'),
		maxLength(255, 'Adresse må maks være 255 tegn'),
	),
	city: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld by'),
		maxLength(255, 'By må maks være 255 tegn'),
	),
	zip: pipe(
		string('Skal være tekst'),
		nonEmpty('Udfyld postnummer'),
		maxLength(255, 'Postnr må maks være 255 tegn'),
	),
	country: pipe(
		string('Skal være tekst'),
		nonEmpty('Vælg land'),
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
		nonEmpty('Beskriv problemet'),
		minLength(50, 'Beskrivelse skal være mindst 50 tegn'),
		maxLength(500, 'Beskrivelse må maks være 500 tegn'),
	),
})

export const load: PageServerLoad = async () => {
	return { form: await superValidate(valibot(schema)) }
}

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const form = await superValidate(request, valibot(schema))
		console.log(form)

		if (!form.valid) return fail(400, { form })

		const response = await fetch('http://localhost:8080/v1/tickets', {
			method: 'post',
			body: JSON.stringify(form.data),
		})

		if (!response.ok) {
			return fail(response.status, { form })
		}

		return message(form, 'ticket created')
	},
}
