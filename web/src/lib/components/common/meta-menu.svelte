<script lang="ts">
	import { goto } from '$app/navigation'
	import * as Command from '$lib/components/ui/command/index.js'
	import { ArrowLeft, Home, List, LogOut, PlusSquare, Search } from '@lucide/svelte'

	let open = $state(false)
	let searchMode = $state(false)
	let searchValue = $state('')

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === ':') {
			e.preventDefault()
			open = !open
		}
		if (e.key === 'Escape' && searchMode) {
			e.preventDefault()
			exitSearchMode()
		}
	}

	function enterSearchMode() {
		searchMode = true
		searchValue = ''
	}

	function exitSearchMode() {
		searchMode = false
		searchValue = ''
	}

	function handleSearchSubmit() {
		if (searchValue.trim()) {
			goto(`/admin/sager/${searchValue.trim().toUpperCase()}`)
			open = false
			exitSearchMode()
		}
	}

	function handleInputChange(value: string) {
		if (searchMode) {
			searchValue = value
		}
	}

	function handleInputKeydown(e: KeyboardEvent) {
		if (searchMode && e.key === 'Enter') {
			e.preventDefault()
			handleSearchSubmit()
		}
	}

	function goBack() {
		history.back()
		open = false
	}

	async function handleLogout() {
		try {
			const response = await fetch('/log-ud?/logout', {
				method: 'POST',
				body: new FormData(),
			})

			if (response.ok) {
				goto('/log-ind')
			}
		} catch (error) {
			console.error('log ud fejlede', error)
		} finally {
			open = false
		}
	}
</script>

<svelte:document onkeydown={handleKeydown} />

<Command.Dialog bind:open loop>
	<Command.Input
		placeholder={searchMode ? 'Indtast RMA sag ID...' : 'Gå til en side eller udfør en handling...'}
		value={searchMode ? searchValue : undefined}
		oninput={e => handleInputChange(e.currentTarget.value)}
		onkeydown={handleInputKeydown}
	/>
	<Command.List>
		{#if searchMode}
			<Command.Group>
				<Command.Item onSelect={handleSearchSubmit}>
					<Search class="size-4" />
					<span>Søg efter {searchValue ? searchValue : 'sag'}</span>
					<Command.Shortcut
						class="rounded-sm border border-b-2 border-l-2 bg-background p-0.5 text-muted-foreground"
						>↵</Command.Shortcut
					>
				</Command.Item>
				<Command.Item onSelect={exitSearchMode}>
					<ArrowLeft />
					<span>Tilbage til menu</span>
				</Command.Item>
			</Command.Group>
		{:else}
			<Command.Group heading="Sider">
				<Command.Item
					onSelect={() => {
						goto('/')
						open = false
					}}
				>
					<Home />
					<span>Forside</span>
				</Command.Item>
				<Command.Item
					onSelect={() => {
						goto('/admin/sager')
						open = false
					}}
				>
					<List />
					<span>Oversigt</span>
				</Command.Item>
				<Command.Item
					onSelect={() => {
						goto('/opret')
						open = false
					}}
				>
					<PlusSquare />
					<span>Opret</span>
				</Command.Item>
			</Command.Group>
			<Command.Separator />
			<Command.Group heading="Handlinger">
				<Command.Item onSelect={enterSearchMode}>
					<Search />
					<span>Søg efter sag</span>
				</Command.Item>
				<Command.Item onSelect={goBack}>
					<ArrowLeft />
					<span>Gå tilbage</span>
				</Command.Item>
				<Command.Item onSelect={handleLogout}>
					<LogOut />
					<span>Log ud</span>
				</Command.Item>
			</Command.Group>
		{/if}
	</Command.List>
</Command.Dialog>
