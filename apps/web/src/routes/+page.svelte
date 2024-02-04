<script lang="ts">
  import { A, Button, Input, Label, P } from 'flowbite-svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import type { Url } from '@repo/server';

  let shortUrl: string;

  const handleSubmit = async (e: SubmitEvent) => {
    const data = new FormData(e.currentTarget as HTMLFormElement);
    const res = await fetch(`${PUBLIC_API_URL}/urls/shorten`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(Object.fromEntries(data))
    });
    const { short_url }: Url = await res.json();
    shortUrl = short_url;
  };
</script>

<form class="grid gap-8 w-full max-w-lg px-3" on:submit|preventDefault={handleSubmit}>
  <div>
    <Label class="block mb-1">Long Url</Label>
    <Input label="url" id="url" name="url" required placeholder="https://www.justinwallace.dev" />
  </div>
  <Button class="w-full" type="submit">poop</Button>
  {#if shortUrl}
    <A href={shortUrl} class="text-center">{shortUrl}</A>
  {/if}
</form>
