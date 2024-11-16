<script lang="ts">
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { sendMessage } from '$lib/services/websocket-service.svelte';

	let response: string;
	function greet(): void {
		const json = JSON.stringify({ appCmd: 'new-cmd', data: { cmdLine: 'echo "Hello World"' } });
		console.log(json);
		sendMessage(json);
	}
</script>

<div class="page-container">
	<h1 class="page-title">Console Page</h1>
	<Input class="my-5 h-20" />
	<!-- <Button on:click={console.log('asd')} type="submit">Execute</Button> -->
	<Button onclick={greet}>Execute</Button>
	<p>Response: {response}</p>
</div>
