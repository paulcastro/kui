/*
 * Copyright 2019-2020 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Registrar, Arguments, WithSourceReferences, Table, isTable } from '@kui-shell/core'

import defaultFlags from './flags'
import { isDryRun, isEntityFormat, KubeOptions, formatOf } from './options'
import { doExecWithStatus, exec } from './exec'
import commandPrefix from '../command-prefix'

import { FinalState } from '../../lib/model/states'
import { isUsage, doHelp } from '../../lib/util/help'
import getSourceRefs from './source'
import { doGetAsEntity, viewTransformer } from './get'

const verbs = ['create', 'apply']

// export const doCreate = (verb: string, command = 'kubectl') => doExecWithStatus(verb, FinalState.OnlineLike, command)

export const doCreate = (verb: string, command = 'kubectl') => async (args: Arguments<KubeOptions>) => {
  if (isUsage(args)) {
    return doHelp(command, args)
  } else {
    if (isDryRun(args)) {
      const raw = await exec(args, undefined, command)
      if (isEntityFormat(formatOf(args))) {
        const entity = await doGetAsEntity(args, raw)
        return entity
      } else {
        return raw.content.stdout
      }
    } else {
      const kuiSourceRef = getSourceRefs(args)
      const table = await doExecWithStatus(verb, FinalState.OnlineLike, command)(args)
      if (isTable(table)) {
        const response: Table & WithSourceReferences = Object.assign({}, table, { kuiSourceRef: await kuiSourceRef })
        return response
      } else {
        return table
      }
    }
  }
}

export const applyFlag = Object.assign({}, defaultFlags, { viewTransformer })

export default (registrar: Registrar) => {
  verbs.forEach(verb => {
    const handler = doCreate(verb)
    registrar.listen(`/${commandPrefix}/kubectl/${verb}`, handler, applyFlag)
    registrar.listen(`/${commandPrefix}/k/${verb}`, handler, applyFlag)
  })
}
