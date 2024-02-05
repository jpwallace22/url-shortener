<script lang="ts">
  import { page } from '$app/stores';
  import { PUBLIC_API_URL } from '$env/static/public';
  import PasswordEye from '$lib/PasswordEye.svelte';
  import { fetchApi } from '$lib/fetch/fetchApi';
  import type { Url } from '@repo/server';
  import { Button, Heading, Input, Label, P } from 'flowbite-svelte';
  import { field, form } from 'svelte-forms';
  import { min, required, url as urlValidator } from 'svelte-forms/validators';

  let showPassword = false;
  const passwordHandler = () => (showPassword = !showPassword);

  const url = field('url', null, [required(), urlValidator()]);
  const password = field('password', null, [min(8)]);
  const myForm = form(url, password);

  const handleSubmit = async () => {
    const id = $page.url.searchParams.get('id');
    const { data, error } = await fetchApi<Url>('/urls/verify', {
      method: 'POST',
      body: JSON.stringify({ password: $password.value, url_id: id })
    });
    if (!error) {
      window.location.href = data.url;
    }
  };
</script>

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

  <Button type="submit" disabled={!$myForm.valid || !$myForm.dirty}>Submit</Button>
</form>
