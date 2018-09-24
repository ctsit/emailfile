# emailfile

Simple command-line program for emailing an attachment via SMTP.


Quickstart
----------

    $ make
    $ env SMTP_HOST=smtp.example.com SMTP_FROM=noreply@example.com \
    > ./emailfile Test tpsreport.pdf boss@example.com <(echo "Attached!")

