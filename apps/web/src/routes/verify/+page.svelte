<script lang="ts">
  import { page } from '$app/stores';
  import { PUBLIC_API_URL } from '$env/static/public';
  import PasswordEye from '../../lib/forms/PasswordEye.svelte';
  import { fetchApi } from '$lib/fetch/fetchApi';
  import type { Url } from '@repo/server';
  import { Button, Heading, Input, Label, P } from 'flowbite-svelte';
  import { field, form } from 'svelte-forms';
  import { min, required, url as urlValidator } from 'svelte-forms/validators';
  import ErrorAlert from '$lib/forms/ErrorAlert.svelte';

  let showPassword = false;
  let errorMessage: string | string[] | null = null;
  const passwordHandler = () => (showPassword = !showPassword);

  const url = field('url', null, [required(), urlValidator()]);
  const password = field('password', null, [min(8)]);
  const verifyForm = form(url, password);

  const handleSubmit = async () => {
    errorMessage = null;
    const id = $page.url.searchParams.get('id');
    const { data, error } = await fetchApi<Url>('/urls/verify', {
      method: 'POST',
      body: JSON.stringify({ password: $password.value, url_id: id })
    });
    if (data) {
      window.location.href = data.url;
      verifyForm.clear();
    }
    if (error) {
      errorMessage = error;
    }
  };
</script>

{#if errorMessage}
  <ErrorAlert message={errorMessage} />
{/if}
<div>
  <Heading class="text-center">Verification Required</Heading>
  <P class="text-center">This link is protected</P>
</div>
<form class="grid gap-8" on:submit|preventDefault={handleSubmit}>
  <div>
    <Label for="password" class="mb-2 inline-flex items-center gap-4">
      Password<PasswordEye {showPassword} {passwordHandler} />
    </Label>
    <Input
      id="password"
      name="password"
      type={showPassword ? 'text' : 'password'}
      placeholder="please enter the password"
      bind:value={$password.value}
    />
  </div>

  <Button type="submit" disabled={!$verifyForm.valid || !$verifyForm.dirty}>Submit</Button>
</form>
