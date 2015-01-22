package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;

import android.view.*;
import android.util.Log;
import android.content.Context;
import android.app.Activity;

import java.lang.reflect.Method;
import java.util.concurrent.Callable;

import java.util.*;

public class RpcHandlerCallViewMethod implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        try {
            String methodName = request.getString("viewMethod");
            String id = request.getString("id");
            JSONArray methodArgs = request.getJSONArray("args");

            View view;
            if (activity.orphanViews.containsKey(Integer.parseInt(id))) {
                view = activity.orphanViews.get(Integer.parseInt(id));
            } else {
                view = ((Activity) context).findViewById(Integer.parseInt(id));
            }

            result = callMethod(
                activity,
                view,
                methodName, request.getString("type"),
                methodArgs
            );
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    protected JSONObject callMethod(
        MainActivity activity,
        final View view, String methodName, String viewType,
        JSONArray methodArgs
    ) throws JSONException {
        JSONObject result = new JSONObject();

        final Object viewObject;

        Method[] allMethods;
        try {
            viewObject = Class.forName(viewType).cast(view);
            allMethods = Class.forName(viewType).getDeclaredMethods();
        } catch(Exception e) {
            // @TODO
            Log.v("!!! CVM methods", e.toString());
            return result;
        }

        final List<Object> requestedParams = new ArrayList<Object>();
        final List<Class> requestedParamTypes = new ArrayList<Class>();

        for (int i = 0; i < methodArgs.length(); i++) {
            try {
                if (methodArgs.get(i) instanceof Integer) {
                    // because get(i) will return Integer class, not int,
                    // which is not acceptable for methods, that wants int.
                    requestedParams.add(methodArgs.getInt(i));
                    requestedParamTypes.add(int.class);
                } else if (methodArgs.get(i) instanceof Double) {
                    // downcast to float
                    requestedParamTypes.add(float.class);
                    requestedParams.add((float) methodArgs.getDouble(i));
                } else {
                    requestedParams.add(methodArgs.get(i));
                    requestedParamTypes.add(methodArgs.get(i).getClass());
                }
            } catch(Exception e) {
                Log.v("!!! CVM fill", e.toString());
            }
        }

        Method targetMethod = null;
        for (Method method : allMethods) {
            if (!method.getName().equals(methodName)) {
                continue;
            }

            Class[] paramTypes = method.getParameterTypes();

            if (paramTypes.length != requestedParams.size()) {
                continue;
            }

            boolean signatureMatched = true;
            int argIndex = 0;
            for (Class paramType : paramTypes) {
                if (!paramType.isAssignableFrom(requestedParamTypes.get(argIndex))) {
                    signatureMatched = false;
                    break;
                }

                argIndex++;
            }

            if (signatureMatched) {
                targetMethod = method;
                break;
            }
        }

        if (targetMethod == null) {
            // @TODO
            return new JSONObject();
        }

        final Method methodToCall = targetMethod;

        try {
            Log.v("!!!", String.format("%s", "before"));
            UIThreadCaller uiThreadCaller = new UIThreadCaller(activity);
            Object callerResult = uiThreadCaller.call(
                new Callable<Object> () {
                    @Override
                    public Object call() throws Exception {
                        Log.v("!!!", String.format("%s", "in"));
                        Object result = methodToCall.invoke(
                            viewObject,
                            requestedParams.toArray()
                        );
                        return result;
                    }
                }
            );

            Log.v("!!!", String.format("%s", "after"));
            JsonConverter converter = new JsonConverter();
            result.put("result", converter.Convert(callerResult));
        } catch (Exception e) {
            result.put("error", e.toString());
        }

        return result;
    }
}
