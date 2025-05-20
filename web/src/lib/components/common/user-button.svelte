<script lang="ts">
	import Button from '../ui/button/button.svelte'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import * as Avatar from '$lib/components/ui/avatar/index.js'
	import type { User } from '$lib/types'
	import { LogOut } from '@lucide/svelte'

	let { user }: { user: User } = $props()
	let initials = $derived(user.name[0] + user.name[user.name.length - 1])
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		{#snippet child({ props })}
			<Button variant="ghost" class="relative size-7 rounded-full" {...props}>
				<Avatar.Root class="size-7 rounded-lg border bg-muted/20">
					<Avatar.Fallback>{initials.toUpperCase()}</Avatar.Fallback>
				</Avatar.Root>
			</Button>
		{/snippet}
	</DropdownMenu.Trigger>
	<DropdownMenu.Content class="w-48" align="end">
		<DropdownMenu.Label class="font-normal">
			<div class="flex flex-col space-y-1">
				<p class="text-sm font-medium leading-none">{user.name}</p>
				<p class="text-xs leading-none text-muted-foreground">{user.email}</p>
			</div>
		</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<form method="POST" action="/log-ud?/logout">
			<button class="w-full">
				<DropdownMenu.Item class="flex items-center justify-between !text-destructive">
					Log ud
					<LogOut class="size-3.5" />
				</DropdownMenu.Item>
			</button>
		</form>
	</DropdownMenu.Content>
</DropdownMenu.Root>
