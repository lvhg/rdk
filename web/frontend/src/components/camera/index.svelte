<script lang="ts">
import type {
  commonApi,
  Client,
  ResponseStream,
  robotApi,
} from '@viamrobotics/sdk';

import Camera from './camera.svelte';
import PCD from '../pcd/index.svelte';
import Collapse from '../collapse.svelte';
import { selectedMap } from '@/lib/camera-state';
import type { StreamManager } from './stream-manager';

export let resources: commonApi.ResourceName.AsObject[];
export let streamManager: StreamManager;
export let client: Client;
export let statusStream: ResponseStream<robotApi.StreamStatusResponse> | null;

const openCameras: Record<string, boolean | undefined> = {};
const refreshFrequency: Record<string, string | undefined> = {};

let triggerRefresh = false;

const setupCamera = (cameraName: string) => {
  openCameras[cameraName] = !openCameras[cameraName];

  for (const camera of resources) {
    if (!refreshFrequency[camera.name]) {
      refreshFrequency[camera.name] = 'Live';
    }
  }
};

const handleRefreshInput = (name: string) => {
  return (event: CustomEvent) => {
    refreshFrequency[name] = event.detail.value;
  };
};

</script>

{#each resources as camera (camera.name)}
  <Collapse title={camera.name}>
    <v-breadcrumbs
      slot="title"
      crumbs="camera"
    />

    <div class="flex flex-col gap-4 border border-t-0 border-medium p-4">
      <v-switch
        label={camera.name}
        aria-label={openCameras[camera.name] ? `Hide Camera: ${camera.name}` : `View Camera: ${camera.name}`}
        value={openCameras[camera.name] ? 'on' : 'off'}
        on:input={() => setupCamera(camera.name)}
      />

      {#if openCameras[camera.name]}
        <div class="flex flex-wrap items-end gap-2">
          <v-select
            value={refreshFrequency[camera.name]}
            class="w-fit"
            label="Refresh frequency"
            aria-label="Refresh frequency"
            options={Object.keys(selectedMap).join(',')}
            on:input={handleRefreshInput(camera.name)}
          />

          {#if refreshFrequency[camera.name] !== 'Live'}
            <v-button
              v-if=""
              icon="refresh"
              label="Refresh"
              on:click={() => {
                triggerRefresh = !triggerRefresh;
              }}
            />
          {/if}
        </div>
      {/if}

      {#if openCameras[camera.name]}
        <Camera
          {client}
          {streamManager}
          {statusStream}
          cameraName={camera.name}
          showExportScreenshot={true}
          refreshRate={refreshFrequency[camera.name]}
        />
      {/if}

      <PCD
        {client}
        cameraName={camera.name}
      />
    </div>
  </Collapse>
{/each}