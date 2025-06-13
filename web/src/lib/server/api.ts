import type { Log, Ticket, TicketWithLogs } from '$lib/types'
import type { NewTicketLog } from '../../routes/(site)/admin/sager/[id]/+page.server'
import type { NewTicket } from '../../routes/(site)/opret/+page.server'
import type { NewPresignedUrl } from '../../routes/api/upload/+server'
import { API_URL } from './env'

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE'

export class ApiError extends Error {
	constructor(
		message: string,
		public status: number,
		public requestID: string,
	) {
		super(message)
		this.name = 'ApiError'
	}
}

const methodsWithBody = new Set<HttpMethod>(['POST', 'PUT', 'PATCH'])

async function apiRequest<TOutput>(
	url: string,
	method: HttpMethod,
	options: {
		body?: unknown
		headers?: Record<string, string>
	} = {},
): Promise<[TOutput, null] | [null, Error]> {
	const { body, headers = {} } = options

	const requestInit: RequestInit = {
		method,
		headers: headers,
	}

	if (methodsWithBody.has(method) && body !== undefined) {
		if (body instanceof FormData) {
			requestInit.body = body
		} else {
			requestInit.body = JSON.stringify(body)
		}
	}

	try {
		const response = await fetch(url, requestInit)

		if (!response.ok) {
			const errorBody = (await response.json()) as { error: string; request_id: string }
			const err = new ApiError(errorBody.error, response.status, errorBody.request_id)
			return [null, err]
		}

		if (response.status === 204 || response.headers.get('content-length') === '0') {
			return [undefined as any, null]
		}

		const contentType = response.headers.get('content-type')

		if (contentType?.includes('application/json')) {
			const json = await response.json()
			return [json, null]
		} else {
			const text = (await response.text()) as any
			return [text, null]
		}
	} catch (error) {
		return [null, error as Error]
	}
}

export const api = {
	async listTickets(token: string, searchparams: URLSearchParams) {
		return apiRequest<{ tickets: Ticket[]; total: number; limit: number }>(
			`${API_URL}/v1/admin/tickets?${searchparams.toString()}`,
			'GET',
			{
				headers: { Authorization: `Bearer ${token}` },
			},
		)
	},
	async listStatuses() {
		return apiRequest<{ statuses: string[] }>(`${API_URL}/v1/tickets/statuses`, 'GET')
	},
	async listCategories() {
		return apiRequest<{ categories: string[] }>(`${API_URL}/v1/tickets/categories`, 'GET')
	},
	async createTicket(data: NewTicket) {
		return apiRequest<{ ticket: TicketWithLogs }>(`${API_URL}/v1/tickets`, 'POST', {
			body: data,
		})
	},
	async getTicket(token: string, id: string) {
		return apiRequest<{ ticket: TicketWithLogs }>(`${API_URL}/v1/admin/tickets/${id}`, 'GET', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
		})
	},
	async createTicketLog(token: string, id: string, data: NewTicketLog) {
		return apiRequest<{ log: Log }>(`${API_URL}/v1/admin/tickets/${id}/logs`, 'POST', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
			body: data,
		})
	},
	async getPublicTicket(id: string) {
		return apiRequest<{ ticket: TicketWithLogs }>(`${API_URL}/v1/tickets/${id}`, 'GET')
	},
	async getPresignedPutUrl(token: string, newPresignedUrl: NewPresignedUrl) {
		return apiRequest<{ presigned_url: string }>(`${API_URL}/v1/admin/upload`, 'PUT', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
			body: newPresignedUrl,
		})
	},
	async updateTicketLog(
		token: string,
		ticketID: string,
		logID: number,
		data: { external_comment: string; internal_comment: string },
	) {
		return apiRequest<undefined>(`${API_URL}/v1/admin/tickets/${ticketID}/logs/${logID}`, 'PUT', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
			body: data,
		})
	},
} as const
