# go-messenger-structs

The package collects the required structs for facebook messenger platform with the supported conversation elements.

# Installation:

```$ go get -u github.com/hellowearemito/go-messenger-structs.git```

### Attachment:

    Defines attachment types. They can be type: template, image, video, audio, file, loacation.

### Element:

    Describes the possible elements of a structure. Title, item url, image url, subtitle, buttons can act as elements.

### Generic Template:

    Identifies generic template. And checking possible errors during the validation. It is possible to add elements to the template.

### List:

    Defines lists with its elements. The facebook list properties.

### Receipt:

    Implements facebooks receipt template.

### Template:

    Handles errors. Like character limit, button limit, bubble limit.

### Action:

    Sender action sending. Mark seen, typing on, typing off.

### Error:

    Defines error structure.

### Events:

    Implements facebook's webhook entry.

### Helper:

    Creates http requests.

###Messagequery:

    Defines content type. Can be test, location. Notification type can be regular, silent push, no push, response, update, message tag, non promotional subscrition.

###Messenger:

    Defines message structure, handler types and debug types.

###Profile:

    Defines profile struct.

###Settings:

    Setting struct defined.

###Thread Control:

    Handles threads. Requests, handlers.