<script lang="ts">
	import '../app.css';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
</script>

<ModeWatcher />
<Sidebar.Provider>
	<AppSidebar />
	<main>
		<Sidebar.Trigger />
		{@render children?.()}
	</main>
</Sidebar.Provider>

<style>
	main {
		@apply flex h-full w-full flex-col;
	}
</style>
