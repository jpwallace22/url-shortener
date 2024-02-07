<script lang="ts">
  import { Toast } from 'flowbite-svelte';
  import { FileCopyOutline } from 'flowbite-svelte-icons';
  import { fade, fly } from 'svelte/transition';

  export let textToCopy: string;
  let status: string | null;

  const copy = (text: string) => {
    try {
      navigator.clipboard.writeText(text);
      status = 'Copied';
    } catch (e) {
      status = 'Error';
    }
  };

  $: if (status) {
    setTimeout(() => (status = null), 1000);
  }
</script>

{#if status}
  <div out:fade={{ duration: 300 }}>
    <Toast dismissable={false} class="absolute -top-10 -left-7 w-20 h-5">{status}</Toast>
  </div>
{/if}
<FileCopyOutline
  {...$$restProps}
  slot="right"
  size="lg"
  class="hover:!text-white cursor-pointer active:text-purple-900"
  on:click={() => copy(textToCopy)}
/>
