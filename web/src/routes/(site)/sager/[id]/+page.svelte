<script lang="ts">
	import ContactInformation from '$lib/components/tickets/contact-information.svelte'
	import Details from '$lib/components/tickets/details.svelte'
	import LogsListExternal from '$lib/components/tickets/logs-list-external.svelte'
	import Title from '$lib/components/tickets/title.svelte'
	import type { Contact } from '$lib/types.js'

	let { data } = $props()
	const ticket = $derived(data.ticket)

	function censor(s: string): string {
		return s.replaceAll(/./g, '*')
	}

	function censorPhone(s: string): string {
		let count = 0
		let result = ''

		for (let i = s.length - 1; i >= 0; i--) {
			const char = s[i]
			if (char !== ' ' && /\d/.test(char)) {
				if (count < 4) {
					result = char + result
					count++
				} else {
					result = '*' + result
				}
			} else {
				result = char + result
			}
		}

		return result
	}

	function censorEmail(email: string): string {
		const [local, domain] = email.split('@')
		if (!domain) return email

		if (local.length <= 2) {
			return '*'.repeat(local.length) + '@' + domain
		}

		const first = local[0]
		const last = local[local.length - 1]
		const stars = '*'.repeat(local.length - 2)
		return `${first}${stars}${last}@${domain}`
	}

	const censoredSender: Contact = $derived.by(() => {
		return {
			company: ticket.sender.company,
			name: censor(ticket.sender.name),
			phone: censorPhone(ticket.sender.phone),
			email: censorEmail(ticket.sender.email),
			street: censor(ticket.sender.street),
			city: ticket.sender.city,
			zip: ticket.sender.zip,
			country: ticket.sender.country,
		}
	})

	const censoredBilling: Contact = $derived.by(() => {
		return {
			company: ticket.billing.company,
			name: censor(ticket.billing.name),
			phone: censorPhone(ticket.billing.phone),
			email: censorEmail(ticket.billing.email),
			street: censor(ticket.billing.street),
			city: ticket.billing.city,
			zip: ticket.billing.zip,
			country: ticket.billing.country,
		}
	})
</script>

<title>RMA Sag #{data.ticket.id} - Skancode RMA Service Portal</title>

<article class="space-y-4">
	<Title {ticket} />
	<div class="flex flex-col gap-4 md:flex-row">
		<Details {ticket} />
		<div class="flex flex-col gap-4 md:w-1/2">
			<ContactInformation contact={censoredSender} title="Afsender" sub="Returaddresse" />
			<ContactInformation
				contact={censoredBilling}
				title="Fakturering"
				sub="Faktureringsaddresse"
			/>
		</div>
	</div>
	<LogsListExternal {ticket} />
</article>
