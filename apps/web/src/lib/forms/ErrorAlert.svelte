<script lang="ts">
  import { Alert, Heading, P } from 'flowbite-svelte';
  import { InfoCircleOutline } from 'flowbite-svelte-icons';

  export let message: string | string[] | null = null;
  const isArray = (text: string | string[] | null): text is string[] => Array.isArray(text);
</script>

<Alert dismissable class="w-full !bg-red-300/20">
  <Heading tag="h4" class="flex gap-3 items-center"><InfoCircleOutline />Uh, Oh.</Heading>
  {#if isArray(message)}
    <ul>
      {#each message as text}
        <li><P>{text}</P></li>
      {/each}
    </ul>
  {:else}
    <P>{message}</P>
  {/if}
</Alert>
