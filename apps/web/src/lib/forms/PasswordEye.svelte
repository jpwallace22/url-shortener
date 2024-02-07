<script lang="ts">
  import { EyeSlashSolid, EyeSolid } from 'flowbite-svelte-icons';

  export let showPassword = false;
  export let passwordHandler: () => void;

  const props = {
    class: 'text-gray-300 hover:text-white'
  };
</script>

<button on:click={passwordHandler} class="w-fit" type="button">
  {#if showPassword}
    <EyeSolid {...props} />
  {:else}
    <EyeSlashSolid {...props} />
  {/if}
</button>
