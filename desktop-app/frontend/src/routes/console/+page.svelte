<script lang="ts">
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button/index.js';

	import { Greet } from '$lib/wailsjs/go/main/App';
	let response: string;
	function greet(): void {
		Greet('Pedro').then((result) => {
			response = result;
		});
	}
</script>

<div class="page-container">
	<h1 class="page-title">Console Page</h1>
	<Input class="my-5 h-20" placeholder="cmd" />
	<!-- <Button on:click={console.log('asd')} type="submit">Execute</Button> -->
	<Button onclick={greet}>Execute</Button>
	<p>Response: {response}</p>
</div>
