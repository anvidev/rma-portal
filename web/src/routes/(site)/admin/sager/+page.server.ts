import { error } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { ApiError } from '$lib/server/api'

export const load: PageServerLoad = async ({ cookies, url, setHeaders, locals }) => {
	const searchParams = url.searchParams
	const token = cookies.get('token') ?? ''

	const [ticketsData, ticketsErr] = await locals.api.listTickets(token, searchParams)
	if (ticketsErr != null) {
		if (ticketsErr instanceof ApiError) {
			return error(ticketsErr.status, {
				message: ticketsErr.message,
				requestId: ticketsErr.requestID,
			})
		} else {
			return error(500, ticketsErr.message)
		}
	}

	const { tickets, total, limit } = ticketsData

	const [statusesData, statusErr] = await locals.api.listStatuses()
	if (statusErr != null) {
		if (statusErr instanceof ApiError) {
			return error(statusErr.status, { message: statusErr.message, requestId: statusErr.requestID })
		} else {
			return error(500, statusErr.message)
		}
	}

	const { statuses } = statusesData

	const [categoriesData, categoriesErr] = await locals.api.listCategories()
	if (categoriesErr != null) {
		if (categoriesErr instanceof ApiError) {
			return error(categoriesErr.status, {
				message: categoriesErr.message,
				requestId: categoriesErr.requestID,
			})
		} else {
			return error(500, categoriesErr.message)
		}
	}

	const { categories } = categoriesData

	setHeaders({
		'Cache-Control': 'private, max-age=120',
	})

	return {
		tickets: tickets,
		statuses: statuses,
		categories: categories,
		total: total,
		limit: limit,
	}
}
