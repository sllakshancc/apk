/*
 * Copyright (c) 2022 WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.apk.apimgt.rest.api.devportal.v1.common.impl;

import org.apache.commons.lang3.StringUtils;
import org.wso2.apk.apimgt.api.APIManagementException;
import org.wso2.apk.apimgt.api.ExceptionCodes;
import org.wso2.apk.apimgt.impl.APIClientGenerationManager;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 *
 * This is the service implementation class for Sdk generate related operations
 */
public class SdkGenCommonImpl {

    private SdkGenCommonImpl() {
    }

    /**
     * Rest API implementation to get the supported sdk languages
     */
    public static List<String> getSdkGenLanguageList() throws APIManagementException {

        APIClientGenerationManager apiClientGenerationManager = new APIClientGenerationManager();
        String supportedLanguages = apiClientGenerationManager.getSupportedSDKLanguages();

        if (StringUtils.isNotEmpty(supportedLanguages)) {
            // Split the string with ',' and add them to a list.
            return Arrays.stream(supportedLanguages.split(",")).collect(Collectors.toList());
        }
        String message = "Could not find the supported sdk languages";
        throw new APIManagementException(message, ExceptionCodes.SDK_NOT_GENERATED);
    }
}
