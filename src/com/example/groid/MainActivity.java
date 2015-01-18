package com.example.groid;

import go.Go;
import go.rpc.Rpc;
import android.app.Activity;
import android.os.Bundle;
import android.content.Context;
import java.util.*;
import org.json.*;
import android.util.Log;

/*
 * MainActivity is the entry point for the groid app.
 *
 * From here, the Go runtime is initialized and a Go function is
 * invoked via gobind language bindings.
 */
public class MainActivity extends Activity {
    protected RpcFrontend mFrontend;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Go.init(getApplicationContext());
        setContentView(R.layout.useless_layout);

        mFrontend = new RpcFrontend(this);
        Rpc.StartBackend(mFrontend);
    }

    public class RpcFrontend extends Rpc.Frontend.Stub {
        protected Context mContext;
        protected Map<String, RpcHandlerInterface> mHandlers;

        RpcFrontend(Context context) {
            mContext = context;
            mHandlers = new HashMap<String, RpcHandlerInterface>();
        }

        public String CallFrontend(final String payload) {
            try {
                JSONObject json = new JSONObject(payload);

                String handlerName = String.format(
                    "%s.RpcHandler%s",
                    this.getClass().getPackage().getName(),
                    json.getString("method")
                );

                // @TODO: handle errors

                RpcHandlerInterface handler = mHandlers.get(handlerName);
                if (handler == null) {
                    handler = (RpcHandlerInterface) Class.forName(
                        handlerName
                    ).newInstance();

                    mHandlers.put(handlerName, handler);
                }

                return handler.Handle(mContext, json).toString();
            } catch (Exception e) {
                // @TODO: proper exception handling
                Log.v("!!!", e.toString());
            }

            // @TODO: properly return json error
            return "error";
        }
    }
}
