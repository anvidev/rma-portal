import { error } from '@sveltejs/kit'
import type { RequestHandler } from './$types'
import { API_URL } from '$lib/server/env'

export const GET: RequestHandler = async ({ url }) => {
	const ticketID = url.searchParams.get('rma')
	if (!ticketID) {
		error(400, 'Kan ikke hente RMA label uden et ID')
	}

	const labelResponse = await fetch(`${API_URL}/v1/tickets/${ticketID}/label`)

	if (!labelResponse.ok) {
		if (labelResponse.status === 404) {
			error(labelResponse.status, 'RMA findes ikke')
		} else {
			error(labelResponse.status, 'Intern server fejl')
		}
	}

	const pdf = await labelResponse.arrayBuffer()

	return new Response(pdf, {
		headers: {
			'Content-Type': 'application/pdf',
			'Content-Disposition': `inline; filename="label-${ticketID}.pdf"`,
		},
	})
}
