import { z } from 'zod'
import type { RequestHandler } from './$types'
import { ApiError } from '$lib/server/api'
import { error } from '@sveltejs/kit'

const schema = z.object({
	key: z.string().nonempty(),
	reference_id: z.string().nonempty().or(z.coerce.number().positive()),
	file_name: z.string().nonempty(),
	file_domain: z.enum(['tickets', 'logs']),
	content_type: z.string().nonempty(),
	content_length: z.coerce.number().positive(),
})

export type NewPresignedUrl = z.infer<typeof schema>

export const PUT: RequestHandler = async ({ request, locals, cookies }) => {
	const token = cookies.get('token')
	if (!token) {
		return new Response('adgang nægtet', { status: 401 })
	}

	const body = await schema.safeParseAsync(await request.json())
	if (!body.success) {
		return new Response('ugyldig data sendt', { status: 400 })
	}

	const [presignedUrlData, err] = await locals.api.getPresignedPutUrl(token, body.data)
	if (err != null) {
		if (err instanceof ApiError) {
			return error(err.status, { message: err.message, requestId: err.requestID })
		} else {
			return error(500, { message: err.message })
		}
	}

	return Response.json({ url: presignedUrlData.presigned_url }, { status: 201 })
}
