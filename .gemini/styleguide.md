# Gemini API Safety Style Guide

## Purpose

This style guide outlines the essential safety considerations for interacting with the Gemini API, specifically focusing on the handling of potentially harmful content.  Adhering to these guidelines ensures that applications built using the Gemini API are responsible and mitigate the risk of generating or displaying dangerous content. The primary focus is on implementing safety settings to block content that falls under the `HARM_CATEGORY_DANGEROUS_CONTENT` category, as well as other harmful categories.


**Requirement:** All media generation requests must include a safety parameter for the `HARM_CATEGORY_DANGEROUS_CONTENT` harm category. It is strongly recommended to also include safety settings for other harm categories, such as hate speech.


```java
    //Example safety builders for harm categories:

     List<SafetySetting> safetySettings = Arrays.asList(
          SafetySetting.newBuilder()
              .setCategory(HarmCategory.HARM_CATEGORY_HATE_SPEECH)
              .setThreshold(SafetySetting.HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE)
              .build(),
          SafetySetting.newBuilder()
              .setCategory(HarmCategory.HARM_CATEGORY_DANGEROUS_CONTENT)
              .setThreshold(SafetySetting.HarmBlockThreshold.BLOCK_MEDIUM_AND_ABOVE)
              .build()
      );

        GenerativeModel model = new GenerativeModel(modelName, vertexAI)
          .withGenerationConfig(generationConfig)
          .withSafetySettings(safetySettings);
```

## Harm Categories and Recommended Thresholds

The Gemini API allows you to configure safety settings to block content based on different harm categories. The following table outlines the available categories and provides recommended thresholds for each.

| Harm Category                     | Description                                                                                                                                                                                                                            | Recommended Threshold                               |
| --------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| `HARM_CATEGORY_HATE_SPEECH`       | Content that promotes discrimination or disparages individuals or groups based on race, ethnicity, religion, gender, sexual orientation, disability, or other protected characteristics.                                                | `BLOCK_MEDIUM_AND_ABOVE`                           |
| `HARM_CATEGORY_DANGEROUS_CONTENT` | Content that promotes, facilitates, or enables harm, or puts people in immediate and serious danger. This includes content that promotes violence, illegal activities, or self-harm.                                                    | `BLOCK_MEDIUM_AND_ABOVE`                           |
| `HARM_CATEGORY_HARASSMENT`        | Content that is abusive, threatening, or intended to bully, intimidate, or harass an individual or group.                                                                                                                                | `BLOCK_MEDIUM_AND_ABOVE` or `BLOCK_LOW_AND_ABOVE` |
| `HARM_CATEGORY_SEXUALLY_EXPLICIT` | Content that is sexually suggestive, contains nudity, or describes sexual acts with the primary intention to cause arousal.                                                                                                              | `BLOCK_MEDIUM_AND_ABOVE` or `BLOCK_LOW_AND_ABOVE` |
| `HARM_CATEGORY_DANGEROUS_CONTENT` | Content that promotes, facilitates, or enables harm, or puts people in immediate and serious danger. This includes content that promotes violence, illegal activities, or self-harm.                                                    | `BLOCK_MEDIUM_AND_ABOVE`                           |
| `HARM_CATEGORY_UNSPECIFIED` | Content that does not fall into any of the other categories, but is still considered harmful.                                                                                                                                                | `BLOCK_NONE` or `BLOCK_LOW_AND_ABOVE` |

**Note:** The `BLOCK_NONE` threshold means that no content will be blocked for that category. `BLOCK_LOW_AND_ABOVE` will block content that is low, medium, or high risk. `BLOCK_MEDIUM_AND_ABOVE` will block content that is medium or high risk. `BLOCK_ONLY_HIGH` will only block content that is high risk.

**Requirement:** All media generation requests must include a safety parameter for the `HARM_CATEGORY_DANGEROUS_CONTENT` harm category. It is strongly recommended to also include safety settings for other harm categories, such as hate speech.
