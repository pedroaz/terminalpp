<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<script lang="ts">
	import { sendMessage } from '$lib/services/websocket-service.svelte';
	import { sendCmdMessage } from '$lib/services/message-factory';
	import { ConsoleState, type Line } from '$lib/services/models';
	import {
		createAndConnectNewConsole as createNewConsole,
		getConsole
	} from '$lib/services/console-manager.svelte';

	let inputElement: HTMLInputElement;
	let id = $state(createNewConsole());

	function handleKeyDown(event: KeyboardEvent): void {
		if (event.key === 'Enter') {
			var json = sendCmdMessage(id, inputElement.value);
			sendMessage(json);
		}
	}

	function focusInput(): void {
		inputElement.focus();
	}
</script>

<div class="console-container" onclick={focusInput}>
	{getConsole(id)?.id}
	{#if getConsole(id)?.state === ConsoleState.LOADING}
		<p>Connecting To Server...</p>
	{:else if getConsole(id)?.state === ConsoleState.READY}
		<div class="flex items-center">
			<div>&nbsp;&nbsp;&gt;</div>
			<input
				bind:this={inputElement}
				type="text"
				class="invisible-input"
				onkeydown={handleKeyDown}
			/>
		</div>
	{:else if getConsole(id)?.state === undefined}
		<p>Creating Console.</p>
	{/if}
	<!-- {#each getConsole(id)?.lines as line}
		<p>-&gt; {line.text}</p>
	{/each} -->
</div>

<style>
	.console-container {
		@apply flex flex-col items-start gap-1 overflow-auto;
	}
	.invisible-input {
		border: none;
		background: none;
		width: 100%;
	}
	.invisible-input:focus {
		outline: none;
		box-shadow: none;
	}
</style>
