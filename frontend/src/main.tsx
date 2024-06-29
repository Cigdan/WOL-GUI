import { StrictMode } from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import { routeTree } from './routeTree.gen';
import Toast from "./components/Toast.tsx";

// Create a new router instance
const router = createRouter({ routeTree });

const queryClient = new QueryClient();

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

// Render the app
const rootElement = document.getElementById('root')!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
      <StrictMode>
        <MantineProvider defaultColorScheme={localStorage.getItem("theme") || "auto"}>
          <QueryClientProvider client={queryClient}>
            <Toast />
            <RouterProvider router={router} />
          </QueryClientProvider>
        </MantineProvider>
      </StrictMode>
  );
}
