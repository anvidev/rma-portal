<script lang="ts" module>
	import { type VariantProps, tv } from 'tailwind-variants'

	export const badgeVariants = tv({
		base: 'focus:ring-ring inline-flex items-center rounded-lg border px-2.5 py-1 text-xs leading-[0.75rem] font-semibold focus:outline-none focus:ring-2 focus:ring-offset-2 w-fit',
		variants: {
			variant: {
				default: 'bg-primary text-primary-foreground border-transparent',
				secondary: 'bg-secondary text-secondary-foreground border-transparent',
				destructive: 'bg-destructive text-destructive-foreground border-transparent',
				outline: 'text-foreground',
			},
		},
		defaultVariants: {
			variant: 'default',
		},
	})

	export type BadgeVariant = VariantProps<typeof badgeVariants>['variant']
</script>

<script lang="ts">
	import type { WithElementRef } from 'bits-ui'
	import type { HTMLAnchorAttributes } from 'svelte/elements'
	import { cn } from '$lib/utils.js'

	let {
		ref = $bindable(null),
		href,
		class: className,
		variant = 'default',
		children,
		...restProps
	}: WithElementRef<HTMLAnchorAttributes> & {
		variant?: BadgeVariant
	} = $props()
</script>

<svelte:element
	this={href ? 'a' : 'span'}
	bind:this={ref}
	{href}
	class={cn(badgeVariants({ variant }), className)}
	{...restProps}
>
	{@render children?.()}
</svelte:element>
