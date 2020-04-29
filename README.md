# Check SSH

## Why

AWS Abuse keep mailing me these days.They warn me about my EC2 Instance has been implicated in activity which resembles attempts to access remote hosts on the internet without authorization.

Then I found some strange IPs connect to my Instance by 22 port.

I have been using the .pem file to connect to SSH,so I could not understand why dose it happend.

## How

1. I replaced port 22;

   * ```shell
     vim /etc/ssh/sshd_config
     ```

   * change listen value

2. Regularly query which ip address tries to connect to SSH port;

   * ```shell
     crontab -e
     ```

   * ```shell
     #Run once every 5 minutes
     5 * * * * YOUPATH/check.sh
     ```

3. Warn me by email.

   * ```shell
     go build YOUPATH/email.go
     ```

## What

When I recevie the email and get some strange IPs:

1. Kill these processes;
2. Put these IPs in the firewall blacklist(like ipset tool);
3. Change the SSH listen port;