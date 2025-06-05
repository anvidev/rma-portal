import { z } from 'zod'
import type { RequestHandler } from './$types'
import { API_URL } from '$lib/server/env'

const schema = z.object({
	key: z.string().nonempty(),
	reference_id: z.string().nonempty().or(z.coerce.number().positive()),
	file_name: z.string().nonempty(),
	file_domain: z.enum(['tickets', 'logs']),
	content_type: z.string().nonempty(),
	content_length: z.coerce.number().positive(),
})

export const PUT: RequestHandler = async ({ request, cookies, fetch }) => {
	const token = cookies.get('token')
	if (!token) {
		return new Response('adgang nægtet', { status: 401 })
	}

	const body = await schema.safeParseAsync(await request.json())
	if (!body.success) {
		console.error(body.error)
		return new Response('ugyldig data sendt', { status: 400 })
	}

	try {
		const response = await fetch(`${API_URL}/v1/admin/upload`, {
			method: 'PUT',
			headers: { Authorization: `Bearer ${token}` },
			body: JSON.stringify(body.data),
		})
		if (!response.ok) {
			return new Response('fejl på api server', {
				status: response.status,
			})
		}

		const { presigned_url } = await response.json()

		return new Response(presigned_url, { status: 201 })
	} catch (error) {
		return new Response('intern server fejl', {
			status: 500,
		})
	}
}
