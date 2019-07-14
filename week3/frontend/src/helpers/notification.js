import { notification } from 'antd';

export const showError = (message) => {
  notification.open({
    message: 'Error',
    description: message
  });
};
