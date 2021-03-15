package com.xhb.logic.http.packet.greet;

import com.xhb.core.packet.HttpPacket;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.greet.model.*;

public class GreetPacket extends HttpPacket<Response> {
	private String name;

	public GreetPacket(String name) {
		super(EmptyRequest.instance);
		
		this.name = name;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.GET;
    }

	@Override
    public String requestUri() {
        return "from/" + name;
    }
}
