using System.Net.Http.Headers;

using HttpClient client = new();
client.DefaultRequestHeaders.Accept.Clear();
client.DefaultRequestHeaders.Accept.Add(
    new MediaTypeWithQualityHeaderValue("application/json"));
client.DefaultRequestHeaders.Add("User-Agent", "csharp-client");

static async Task SendMessageAsync(HttpClient client, Proto.SayHelloRequest request)
{
    HttpContent content = new StringContent(request.ToString());
    
    var json = await client.PostAsync(
        "http://localhost:8080/v1/echo",
        content
    );

    Console.Write(json);
}


Proto.SayHelloRequest sayHelloRequest = new Proto.SayHelloRequest();

long timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();
Proto.Payload firstMessage = new Proto.Payload { Source = "csharp-client", SourceId = "123-321", SourceTimestamp = timestamp };
firstMessage.Payload_.Add("userID", "987-789");

sayHelloRequest.Data.Add(firstMessage);

await SendMessageAsync(client, sayHelloRequest);
