<script lang="ts">
	import '../app.css';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
	import { connectWebSocket } from '$lib/services/websocket-service.svelte';
	connectWebSocket();
</script>

<ModeWatcher />
<Sidebar.Provider>
	<AppSidebar />
	<main>
		<Sidebar.Trigger />
		<div>
			{@render children?.()}
		</div>
	</main>
</Sidebar.Provider>

<style>
	main {
		width: 100%;
		height: 100vh;
		/* background-color: red; */
	}
</style>
