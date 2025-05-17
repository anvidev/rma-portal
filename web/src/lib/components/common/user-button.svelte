<script lang="ts">
	import Button from '../ui/button/button.svelte'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import * as Avatar from '$lib/components/ui/avatar/index.js'
	import type { User } from '$lib/types'

	let { user }: { user: User } = $props()
	let initials = $derived(user.name[0] + user.name[user.name.length - 1])
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		{#snippet child({ props })}
			<Button variant="ghost" class="relative h-8 w-8 rounded-full" {...props}>
				<Avatar.Root class="bg-muted/20 size-8 rounded-lg border">
					<Avatar.Fallback>{initials.toUpperCase()}</Avatar.Fallback>
				</Avatar.Root>
			</Button>
		{/snippet}
	</DropdownMenu.Trigger>
	<DropdownMenu.Content class="w-56" align="end">
		<DropdownMenu.Label class="font-normal">
			<div class="flex flex-col space-y-1">
				<p class="text-sm font-medium leading-none">{user.name}</p>
				<p class="text-muted-foreground text-xs leading-none">{user.email}</p>
			</div>
		</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<form method="POST" action="/log-ud?/logout">
			<button class="w-full">
				<DropdownMenu.Item class="!text-destructive">
					Log out
					<DropdownMenu.Shortcut>⇧⌘Q</DropdownMenu.Shortcut>
				</DropdownMenu.Item>
			</button>
		</form>
	</DropdownMenu.Content>
</DropdownMenu.Root>
