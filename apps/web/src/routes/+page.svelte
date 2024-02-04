<script lang="ts">
  import { PUBLIC_API_URL } from '$env/static/public';
  import CopyButton from '$lib/CopyButton.svelte';
  import PasswordEye from '$lib/PasswordEye.svelte';
  import type { Url } from '@repo/server';
  import { Button, Checkbox, Heading, Input, Label, P } from 'flowbite-svelte';
  import { field, form } from 'svelte-forms';
  import { min, required, url as urlValidator } from 'svelte-forms/validators';
  import { fade, fly } from 'svelte/transition';

  let shortUrl: string;
  let hasPassword = false;
  let showPassword = false;

  const url = field('url', null, [required(), urlValidator()]);
  const password = field('password', null, [min(8)]);
  const myForm = form(url, password);

  const passwordHandler = () => (showPassword = !showPassword);

  const handleSubmit = async (e: SubmitEvent) => {
    const res = await fetch(`${PUBLIC_API_URL}/urls/shorten`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify({ url: $url.value, password: $password.value })
    });
    const { short_url }: Url = await res.json();
    shortUrl = short_url;
  };
</script>

<div class="w-full grid gap-16 max-w-lg px-3">
  <div>
    <Heading class="text-center">Link Shrinker</Heading>
    <P class="text-center">The shorter the sweeter</P>
  </div>
  <form class="grid gap-8" on:submit|preventDefault={handleSubmit}>
    <div class="grid gap-4">
      <div>
        <div>
          <Label for="url" class="mb-2">URL to shorten</Label>
          <Input
            id="url"
            name="url"
            placeholder="https://mycoolwebsite.com"
            bind:value={$url.value}
          />
        </div>
      </div>
      {#if hasPassword}
        <div in:fade>
          <div in:fly={{ x: -300 }}>
            <Label for="password" class="mb-2 inline-flex items-center gap-4"
              >Password <PasswordEye {showPassword} {passwordHandler} /></Label
            >
            <Input
              id="password"
              name="password"
              type={showPassword ? 'text' : 'password'}
              placeholder="8 character minimum"
              bind:value={$password.value}
            />
          </div>
        </div>
      {/if}
      <Checkbox bind:checked={hasPassword}>Password Protected</Checkbox>
      <Button type="submit" disabled={!$myForm.valid || !$myForm.dirty}>Make it smaller!</Button>
    </div>
    {#if shortUrl}
      <Input class="!bg-transparent" value={shortUrl} readonly>
        <CopyButton slot="right" textToCopy={shortUrl} />
      </Input>
    {/if}
  </form>
</div>
