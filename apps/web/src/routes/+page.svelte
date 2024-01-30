<script lang="ts">
	import { MyCounterButton } from '@repo/ui';

	export let data;
</script>

<h1>Web</h1>
<MyCounterButton />
<h1>{data.message}</h1>
