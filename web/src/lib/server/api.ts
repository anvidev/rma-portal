import type { Ticket, TicketWithLogs } from '$lib/types'
import type { NewTicketLog } from '../../routes/(site)/admin/sager/[id]/+page.server'
import type { NewTicket } from '../../routes/(site)/opret/+page.server'
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
		headers: {
			'Content-Type': 'application/json',
			...headers,
		},
	}

	// TODO: support formdata also???
	if (methodsWithBody.has(method) && body !== undefined) {
		requestInit.body = JSON.stringify(body)
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
		return apiRequest<{ categories: string[] }>(`${API_URL}/v1/tickets/statuses`, 'GET')
	},
	async createTicket(data: NewTicket) {
		return apiRequest<{ ticket: TicketWithLogs }>(`${API_URL}/v1/tickets`, 'POST', {
			body: data,
		})
	},
	async getTicket(token: string, id: string) {
		return apiRequest<{ ticket: Ticket }>(`${API_URL}/v1/admin/tickets/${id}`, 'GET', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
		})
	},
	async createTicketLog(token: string, id: string, data: NewTicketLog) {
		return apiRequest<null>(`${API_URL}/v1/admin/tickets/${id}/log`, 'POST', {
			headers: {
				Authorization: `Bearer ${token}`,
			},
			body: data,
		})
	},
	async getPublicTicket(id: string) {
		return apiRequest<{ ticket: TicketWithLogs }>(`${API_URL}/v1/tickets/${id}`, 'GET')
	},
} as const
