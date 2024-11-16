<script lang="ts">
	import House from 'lucide-svelte/icons/house';
	import SquareChartGantt from 'lucide-svelte/icons/square-chart-gantt';
	import SquareChevronRight from 'lucide-svelte/icons/square-chevron-right';
	import Settings from 'lucide-svelte/icons/settings';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { isWebsocketConnected } from '$lib/services/websocket-service.svelte';
	// Menu items.
	const items = [
		{
			title: 'Home',
			url: '/',
			icon: House
		},
		{
			title: 'Projects',
			url: 'projects',
			icon: SquareChartGantt
		},
		{
			title: 'Console',
			url: 'console',
			icon: SquareChevronRight
		},
		{
			title: 'Settings',
			url: 'settings',
			icon: Settings
		}
	];
</script>

<Sidebar.Root>
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel
				><div class="flex flex-auto items-center justify-around gap-5">
					<h1>Console++</h1>
				</div></Sidebar.GroupLabel
			>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each items as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
	<Sidebar.Footer
		>Socket Connection: {isWebsocketConnected() ? 'Connected' : 'Disconnected'}</Sidebar.Footer
	>
</Sidebar.Root>
