<script lang="ts">
	let { data } = $props()
</script>

<h1>{data.ticket?.sender.name}</h1>
<h1>{data.ticket?.sender.email}</h1>
