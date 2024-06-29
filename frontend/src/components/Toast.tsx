import { Toaster, ToastBar } from 'react-hot-toast';
import { Notification } from '@mantine/core';

const Toast = () => (
    <Toaster
        position={"bottom-right"}
        toastOptions={{
          success: {
            loading: false,
            type: 'success',
          },
          error: {
            loading: false,
            type: 'error',
          },
          duration: 5000,
        }}
    >
      {(t) => (
          <ToastBar
              style={{
                padding: 0,
                margin: 0,
                backgroundColor: "transparent",
              }}
              toast={t}
          >
            {({ message }) => {
              let color;
              let title;
              switch(t.type) {
                case 'success':
                  color = 'green';
                  title = "Success!"
                  break;
                case 'error':
                  color = 'red';
                  title = "Error!"
                  break;
                default:
                  color = 'default'; // use default color for loading or other types
              }

              return (
                  <Notification
                      color={color}
                      title={title}
                      withBorder
                  >
                    {message}
                  </Notification>
              );
            }}
          </ToastBar>
      )}
    </Toaster>
);

export default Toast;
